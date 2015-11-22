# gocodeit

Utility to manage [gocode][1]'s lib-path when using [GO15VENDOREXPERIMENT][2] with [glide][3].

## Why?

As [gocode][1] is currently unaware of [GO15VENDOREXPERIMENT][2], it doesn't automatically use your project's vendor directory, so you don't have autocompletion in your editor for vendorized packages.

But [gocode][1] supply a command ($ gocode set lib-path) for you to view and set a path[s] to its lib-path, and you can add your project's vendor folder to it.

Great, but it's a PITA to manually manage this, so I created this tool to help with it:

    $ gocodeit set
    Current project's vendor successfully added to gocode's lib-path

You can remove it with:

    $ gocodeit unset
    Current project's vendor successfully removed to gocode's lib-path

Check the help command for the other things you can do with it.

## Note

Remember to run rebuild your dependencies every time you change your vendor packages:

    $ glide rebuild

## Contributing

1. Fork it
2. Checkout the develop branch (`git checkout -b develop`)
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Commit your changes (`git commit -am 'Add some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create new Pull Request

## License

This application is distributed under the [MIT License][4].

[1]: https://github.com/nsf/gocode
[2]: https://golang.org/s/go15vendor
[3]: https://github.com/Masterminds/glide
[4]: https://en.wikipedia.org/wiki/MIT_License
