package brandit

import (
	"testing"

	"github.com/sonavilabs/lego/v4/platform/tester"
	"github.com/stretchr/testify/require"
)

const envDomain = envNamespace + "DOMAIN"

var envTest = tester.NewEnvTest(EnvAPIKey, EnvAPIUsername).WithDomain(envDomain)

func TestNewDNSProvider(t *testing.T) {
	testCases := []struct {
		desc     string
		envVars  map[string]string
		expected string
	}{
		{
			desc: "success",
			envVars: map[string]string{
				EnvAPIKey:      "key",
				EnvAPIUsername: "test_user",
			},
		},
		{
			desc: "missing API key",
			envVars: map[string]string{
				EnvAPIUsername: "test_user",
			},
			expected: "brandit: some credentials information are missing: BRANDIT_API_KEY",
		},
		{
			desc: "missing secret",
			envVars: map[string]string{
				EnvAPIKey: "key",
			},
			expected: "brandit: some credentials information are missing: BRANDIT_API_USERNAME",
		},
		{
			desc:     "missing credentials",
			envVars:  map[string]string{},
			expected: "brandit: some credentials information are missing: BRANDIT_API_KEY,BRANDIT_API_USERNAME",
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
				require.NotNil(t, p.client)
			} else {
				require.EqualError(t, err, test.expected)
			}
		})
	}
}

func TestNewDNSProviderConfig(t *testing.T) {
	testCases := []struct {
		desc     string
		apiKey   string
		user     string
		expected string
	}{
		{
			desc:   "success",
			apiKey: "key",
			user:   "test_user",
		},
		{
			desc:     "missing API key",
			user:     "test_user",
			expected: "brandit: credentials missing",
		},
		{
			desc:     "missing secret",
			apiKey:   "key",
			expected: "brandit: credentials missing",
		},
		{
			desc:     "missing credentials",
			expected: "brandit: credentials missing",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			config := NewDefaultConfig()
			config.APIKey = test.apiKey
			config.APIUsername = test.user

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

	err = provider.CleanUp(envTest.GetDomain(), "", "123d==")
	require.NoError(t, err)
}
