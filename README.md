# Latin Macronizer CLI

This is a simple command line tool that makes it easy to do single word-form lookups. It uses the
macrons data generated from [Alatius/latin-macronizer][latin-macronizer].

[latin-macronizer]: https://github.com/Alatius/latin-macronizer

## Usage

The program will present you with a simple prompt. For any text that is entered, a list of
matching word forms is returned.

By default any word forms starting with the entered text are returned. To restrict the results to
exact matches only, add a space at the end of the text.

Note that there is a limit to the number of entries returned for any given query.

## Development

Run `make` to build the `mzcli` command.

The macrons data is first pre-processed to put it in more compact form. Then it is included as a
static, gzipped resource in the compiled binary. See the `Makefile` for details.
