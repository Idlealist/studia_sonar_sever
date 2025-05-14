# studia_sonar_sever

.git/hooks/pre-commit:

` 
#!/bin/sh

golangci-lint run
RESULT=$?

if [ $RESULT -ne 0 ]; then
    echo "Go linter founs issues aborting commit"
    exit 1
fi

echo "Lint passed, proceeding with the commit"
exit 0
`