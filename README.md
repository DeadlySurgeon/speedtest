# Speedtest

Just a wrapper around the speedtest cli offered up by speedtest.net

This implementation is for a friend help complete a basic project, I don't
recommend using this tool.

## Speedtest CLI

**Tested Version**: 1.0.0.2

The speedtest CLI _must_ be present in the path. To verify that you have it
installed correctly, you may run the command `speedtest --version` and expect to
see an output similar to this:

```
Speedtest by Ookla 1.0.0.2 (5ae238b) Darwin 19.3.0 x86_64

The official command line client for testing the speed and performance
of your internet connection.
```

Instructions on how to install the CLI can be found
[here](https://www.speedtest.net/apps/cli) and a blog post by speedtest.net
talking about this tool can be found
[here](https://www.speedtest.net/insights/blog/introducing-speedtest-cli/).

## Example:

This code blocks until it completes the test. 

```golang
result, err := speedtest.NewTest()
```

## Contributing

If anyone decides to use this, any contribution _must_ reach a 100% test
coverage. There might be a few exceptions in the future, but realistically I
don't see a reason for a wrapper, or a code set this small for that matter, to
not be covered completely.

All code should pass through the bog standard gofmt. Any PRs that don't meet
formatting standards will be rejected.

## License

```
    Copyright 2020 deadly.surgery

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use these files except in compliance with the License.
    You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
```
