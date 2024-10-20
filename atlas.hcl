data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./cmd/gorm-loader",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = getenv("DATABASE_DSN")
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}