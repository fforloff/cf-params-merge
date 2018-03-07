## CFmingle ##

CFmingle is an AWS CloudFormation helper CLI tool. It allows merging parameter values from different sources.
It takes CloudFormation template files in JSON or YAML formats

In order or preference:
* Unix shell environment variables with the same name as parameters in a given ClopudFormation template
* one or more paramter files - as per [AWS Documentation](https://aws.amazon.com/blogs/devops/passing-parameters-to-cloudformation-stacks-with-the-aws-cli-and-powershell/), 
  where the precedence increases from the first to the last parameter file.

## Testing ##
```
make test
```

## Compiling ##
```
make
```

## Usage ##
```
Usage:
  cfmingle [command]

Available Commands:
  help        Help about any command
  merge       Merge parameters

Flags:
  -h, --help   help for cfmingle

Use "cfmingle [command] --help" for more information about a command.
```
```
cfmingle merge --help 
Merge Cloudformation parameters from parameters files and environment variables

Usage:
  cfmingle merge [flags]

Flags:
  -h, --help                     help for merge
  -p, --param-file stringArray   CloudFormation parameter file
  -t, --template string          CloudFormation template (Required)
```

## Example with a JSON template ##
```
Parameter3="55" ./cfmingle merge -t test/template.json -p test/params1.json  -p test/params2.json | jq
[
  {
    "ParameterKey": "Parameter1",
    "ParameterValue": "parameter 1 value"
  },
  {
    "ParameterKey": "Parameter3",
    "ParameterValue": "55"
  },
  {
    "ParameterKey": "Parameter4",
    "UsePreviousValue": true
  }
]
```
## Example with a YAML template ##
```
Parameter3="55" ./cfmingle merge -t test/template.yaml -p test/params1.json  -p test/params2.json | jq
[
  {
    "ParameterKey": "Parameter1",
    "ParameterValue": "parameter 1 value"
  },
  {
    "ParameterKey": "Parameter3",
    "ParameterValue": "55"
  },
  {
    "ParameterKey": "Parameter4",
    "UsePreviousValue": true
  }
]
```