
# !/bin/bash
printf "\n regenerated gqlLen files \n"

rm -f pkg/gql/exec/tmp/generated.go \
    pkg/gql/models/tmp/generated.go \
    pkg/gql/resolvers/tmp/generated.go \ 
time go run -v github.com/99designs/gqlgen $1
printf "\n Done! \n"