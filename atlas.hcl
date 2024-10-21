data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./cmd/gorm-loader",
  ]
}

data "external" "config" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./cmd/gorm-loader/config/main.go",
  ]
}

locals{
    config = jsondecode(data.external.config)
}

env "gorm" {
  src = data.external_schema.gorm.url
  url = local.config.databaseUrl
  dev = "docker://postgres/14"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}