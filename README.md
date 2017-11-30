# We're waiting

It's a bit boring to wait for a long script to complete,
so it'll be better if someone will wait with you!

[![gif with examples][https://raw.githubusercontent.com/nvbn/we-are-waiting/master/example.gif]][https://youtu.be/QRb5l8AF2O0]

*I don't know why the gif is so shitty*

## Installation

```bash
sudo curl -L https://github.com/nvbn/we-are-waiting/releases/download/0.1/waiting-`uname -s`-`uname -m` -o /usr/local/bin/waiting
```

## Usage

Wait for a long script:

```bash
./long-script.sh | waiting
```

You cna change the speed of *waiting* with `--tick` argument.

## License MIT
