laut-live
=========

A small command line tool to stream STDIN to a laut.fm icecast station.

Synopsis
--------

Download the Linux binary:

```
curl https://raw.githubusercontent.com/niko/lautfm-streamer/master/lautfm-streamer.linux.64 > lautfm-streamer
chmod 755 lautfm-streamer
```

Use `curl` to fetch another stream and redirect to you laut.fm station: 

```
curl -L http://stream.laut.fm/eins | ./lautfm-streamer yourstation xxxx-xxxx-xxxx-xxxx
curl -L http://stream.laut.fm/eins | ./lautfm-streamer yourstation xxxx-xxxx-xxxx-xxxx metadata
```

The optional third argument `metadata` is a file containing the metadata sent to the server. The file is watched for changes so updates are promoted.

Stream a local mp3:

```
cat Atmosphere/WhenLiveGivesYouLemons/02-Puppets.mp3 | ./laut-live yourstation xxxx-xxxx-xxxx-xxxx
```

You can easily use lame to transcode a file on the fly:

```
lame ~/Downloads/Atmosphere/WhenLiveGivesYouLemons/02-Puppets.mp3 - | ./laut-live yourstation xxxx-xxxx-xxxx-xxxx
```

As lame defaults to 128kbit/44.1khz this works without further parameters. Transcode a stream on the fly:

```
curl -L http://stream.laut.fm/eins | lame - - | ./laut-live yourstation xxxx-xxxx-xxxx-xxxx
```

You can add auth parameters to curl to stream password protected streams. Or add the `--max-time` option to automatically stop after a given number of seconds.

Requirements
------------

libshout must be installed.

Restrictions
-----

As cross compilation does't work for GO projects with C bindings I can only provide the linux 64 bit binary.

