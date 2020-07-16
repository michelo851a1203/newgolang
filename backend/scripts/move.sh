printf "\n regenerated gqlLen files \n"

cp pkg/gql/exec/tmp/generated.go pkg/gql \
&& mv pkg/gql/generated.go pkg/gql/main.go \

cp pkg/gql/models/tmp/generated.go pkg/gql/models \
&& mv pkg/gql/models/generated.go pkg/gql/models/models.go \

cp pkg/gql/resolvers/tmp/generated.go pkg/gql/resolvers \
&& mv pkg/gql/resolvers/generated.go pkg/gql/resolvers/main.go

printf "\n Done! \n"