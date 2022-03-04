module naka-disc/private-standard-layout

go 1.17

require internal/sample v1.0.0

replace internal/sample => ./internal/app/sample

require internal/hello v1.0.0

replace internal/hello => ./internal/app/hello
