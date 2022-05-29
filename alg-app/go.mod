module github.com/AllenShi/generic-modules/alg-app

go 1.18

require (
  "github.com/AllenShi/generic-modules/alg-lib/math" v0.0.0
)

replace (
  github.com/AllenShi/generic-modules/alg-lib => "../alg-lib"
)
