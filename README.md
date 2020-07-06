# Latin Macronizer CLI

This is a simple command line tool that makes it easy to do single word-form lookups. It uses the
macrons data generated from [Alatius/latin-macronizer][latin-macronizer].

[latin-macronizer]: https://github.com/Alatius/latin-macronizer

## Development

You will need to have [binclude] installed.

Run `make` to build the `mzcli` command.

The macrons data is first pre-processed to put it in more compact form. Then it is included as a
static, gzipped resource in the compiled binary. See the `Makefile` for details.

[binclude]: https://github.com/lu4p/binclude
