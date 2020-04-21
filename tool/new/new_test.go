package new

import (
	"os"
	"testing"
	"github.com/jukylin/esim/log"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/jukylin/esim/pkg/file-dir"
)


func TestProject_Run(t *testing.T) {
	project := NewProject(log.NewLogger())

	v := viper.New()

	v.Set("server_name", "example-a")
	v.Set("gin", true)

	project.Run(v)

	exists, err := file_dir.IsExistsDir("example-a")
	assert.Nil(t, err)
	if exists {
		os.RemoveAll("example-a")
	}
}

func TestProject_GetPackName(t *testing.T) {
	project := NewProject(log.NewNullLogger())

	testCases := []struct{
		caseName string
		serverName string
		expected string
	}{
		{"case1", "api-test", "api_test"},
		{"case2", "api-test-user", "api_test_user"},
		{"case3", "test", "test"},
	}

	for _, test := range testCases{
		t.Run(test.caseName, func(t *testing.T) {
			project.ServerName = test.serverName
			project.getPackName()
			assert.Equal(t, test.expected, project.PackageName)
		})
	}
}

func TestProject_CheckServiceName(t *testing.T) {
	project := NewProject(log.NewNullLogger())

	testCases := []struct{
		caseName string
		serviceName string
		expected bool
	}{
		{"case1", "api_test", true},
		{"case2", "api1123", false},
		{"case3", "example&*^", false},
		{"case4", "api-test", true},
	}

	for _, test := range testCases{
		t.Run(test.caseName, func(t *testing.T) {
			project.ServerName = test.serviceName
			result := project.checkServerName()
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestProject_BindInput(t *testing.T) {
	project := NewProject(log.NewNullLogger())

	v := viper.New()
	v.Set("service_name", "example")

	project.bindInput(v)
}

func TestProject_ExecuteTmpl(t *testing.T) {
	project := NewProject(log.NewNullLogger())
	project.ServerName = "test"

	text := `/{{.ServerName}}`

	result, err := project.executeTmpl("test", text)
	assert.Nil(t, err)
	assert.Equal(t, "/test", string(result))
}