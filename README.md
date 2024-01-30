# go templ starter

This is a demo project to show how to create a full stack web application using Go, Echo templ, HTMX, Supabase, Postgres and Tailwind CSS.

We use the following tools:

| Category       | Technology                                                                 |
| -------------- | -------------------------------------------------------------------------- |
| Language       | [Go](https://golang.org/)                                                  |
| Web Framework  | [Echo](https://echo.labstack.com/)                                        |
| Templating     | [templ](https://github.com/a-h/templ)                                      |
| JavaScript     | [HTMX](https://htmx.org/)                                                  |
| CSS            | [Tailwind CSS](https://tailwindcss.com/)                                   |
| Database/Auth  | [Supabase](https://supabase.com/), [PostgreSQL](https://www.postgresql.org/) |
| Build Tools    | [Docker](https://www.docker.com/), [Make](https://www.gnu.org/software/make/), [GitHub Actions](https://docs.github.com/en/actions) |

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/)
- [Make](https://www.gnu.org/software/make/)
- [Go](https://golang.org/)


### Installation

1. Clone the repo

```sh
   git clone https://github.com/aaron-smits/go-templ-starter.git
```

2. Set environment variables

```sh
   cp .env.dev.example .env
```

3. Start the database and the app with docker compose

```sh
   make docker
```

## todo

### p0

- [ ] Update frontend styling
  - [ ] Header/sidebar
  - [ ] Footer
  - [ ] Space between todo cards
  - [ ] Light/dark modes

### p1

- [ ] Add a `seed` command to the Makefile to seed the dev database (postgres docker container)

### p2

- [ ] Write e2e tests
- [ ] Write unit tests
- [ ] Add a `test` command to the Makefile

### p3
