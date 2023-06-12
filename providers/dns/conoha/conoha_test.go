package conoha

import (
	"testing"
	"time"

	"github.com/sonavilabs/lego/v4/platform/tester"
	"github.com/stretchr/testify/require"
)

const envDomain = envNamespace + "DOMAIN"

var envTest = tester.NewEnvTest(
	EnvTenantID,
	EnvAPIUsername,
	EnvAPIPassword).
	WithDomain(envDomain)

func TestNewDNSProvider(t *testing.T) {
	testCases := []struct {
		desc     string
		envVars  map[string]string
		expected string
	}{
		{
			desc: "complete credentials, but login failed",
			envVars: map[string]string{
				EnvTenantID:    "tenant_id",
				EnvAPIUsername: "api_username",
				EnvAPIPassword: "api_password",
			},
			expected: `conoha: failed to login: unexpected status code: [status code: 401] body: {"unauthorized":{"message":"Invalid user: api_username","code":401}}`,
		},
		{
			desc: "missing credentials",
			envVars: map[string]string{
				EnvTenantID:    "",
				EnvAPIUsername: "",
				EnvAPIPassword: "",
			},
			expected: "conoha: some credentials information are missing: CONOHA_TENANT_ID,CONOHA_API_USERNAME,CONOHA_API_PASSWORD",
		},
		{
			desc: "missing tenant id",
			envVars: map[string]string{
				EnvTenantID:    "",
				EnvAPIUsername: "api_username",
				EnvAPIPassword: "api_password",
			},
			expected: "conoha: some credentials information are missing: CONOHA_TENANT_ID",
		},
		{
			desc: "missing api username",
			envVars: map[string]string{
				EnvTenantID:    "tenant_id",
				EnvAPIUsername: "",
				EnvAPIPassword: "api_password",
			},
			expected: "conoha: some credentials information are missing: CONOHA_API_USERNAME",
		},
		{
			desc: "missing api password",
			envVars: map[string]string{
				EnvTenantID:    "tenant_id",
				EnvAPIUsername: "api_username",
				EnvAPIPassword: "",
			},
			expected: "conoha: some credentials information are missing: CONOHA_API_PASSWORD",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			defer envTest.RestoreEnv()
			envTest.ClearEnv()

			envTest.Apply(test.envVars)

			p, err := NewDNSProvider()

			if test.expected == "" {
				require.NoError(t, err)
				require.NotNil(t, p)
				require.NotNil(t, p.config)
			} else {
				require.EqualError(t, err, test.expected)
			}
		})
	}
}

func TestNewDNSProviderConfig(t *testing.T) {
	testCases := []struct {
		desc     string
		expected string
		tenant   string
		username string
		password string
	}{
		{
			desc:     "complete credentials, but login failed",
			expected: `conoha: failed to login: unexpected status code: [status code: 401] body: {"unauthorized":{"message":"Invalid user: api_username","code":401}}`,
			tenant:   "tenant_id",
			username: "api_username",
			password: "api_password",
		},
		{
			desc:     "missing credentials",
			expected: "conoha: some credentials information are missing",
		},
		{
			desc:     "missing tenant id",
			expected: "conoha: some credentials information are missing",
			username: "api_username",
			password: "api_password",
		},
		{
			desc:     "missing api username",
			expected: "conoha: some credentials information are missing",
			tenant:   "tenant_id",
			password: "api_password",
		},
		{
			desc:     "missing api password",
			expected: "conoha: some credentials information are missing",
			tenant:   "tenant_id",
			username: "api_username",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			config := NewDefaultConfig()
			config.TenantID = test.tenant
			config.Username = test.username
			config.Password = test.password

			p, err := NewDNSProviderConfig(config)

			if test.expected == "" {
				require.NoError(t, err)
				require.NotNil(t, p)
				require.NotNil(t, p.config)
				require.NotNil(t, p.client)
			} else {
				require.EqualError(t, err, test.expected)
			}
		})
	}
}

func TestLivePresent(t *testing.T) {
	if !envTest.IsLiveTest() {
		t.Skip("skipping live test")
	}

	envTest.RestoreEnv()
	provider, err := NewDNSProvider()
	require.NoError(t, err)

	err = provider.Present(envTest.GetDomain(), "", "123d==")
	require.NoError(t, err)
}

func TestLiveCleanUp(t *testing.T) {
	if !envTest.IsLiveTest() {
		t.Skip("skipping live test")
	}

	envTest.RestoreEnv()
	provider, err := NewDNSProvider()
	require.NoError(t, err)

	time.Sleep(1 * time.Second)

	err = provider.CleanUp(envTest.GetDomain(), "", "123d==")
	require.NoError(t, err)
}
