# studia_sonar_sever

.git/hooks/pre-commit:

``` bash
#!/bin/sh

golangci-lint run
RESULT=$?

if [ $RESULT -ne 0 ]; then
    echo "Go linter founs issues aborting commit"
    exit 1
fi

echo "Lint passed, proceeding with the commit"
exit 0
```

[![SonarQube Cloud](https://sonarcloud.io/images/project_badges/sonarcloud-highlight.svg)](https://sonarcloud.io/summary/new_code?id=Idlealist_studia_sonar_sever)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=Idlealist_studia_sonar_sever&metric=bugs)](https://sonarcloud.io/summary/new_code?id=Idlealist_studia_sonar_sever)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=Idlealist_studia_sonar_sever&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=Idlealist_studia_sonar_sever)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=Idlealist_studia_sonar_sever&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=Idlealist_studia_sonar_sever)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=Idlealist_studia_sonar_sever&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=Idlealist_studia_sonar_sever)