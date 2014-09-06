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
cat Atmosphere/WhenLiveGivesYouLemons/02-Puppets.mp3 | ./lautfm-streamer yourstation xxxx-xxxx-xxxx-xxxx
```

You can easily use lame to transcode a file on the fly:

```
lame ~/Downloads/Atmosphere/WhenLiveGivesYouLemons/02-Puppets.mp3 - | ./lautfm-streamer yourstation xxxx-xxxx-xxxx-xxxx
```

As lame defaults to 128kbit/44.1khz this works without further parameters. Transcode a stream on the fly:

```
curl -L http://stream.laut.fm/eins | lame - - | ./lautfm-streamer yourstation xxxx-xxxx-xxxx-xxxx
```

You can add auth parameters to curl to stream password protected streams. Or add the `--max-time` option to automatically stop after a given number of seconds.

You can use a song recognition system like gracenote to add metadata automatically:

```
mkfifo pipe
curl http://stream1.laut.fm/eins | tee pipe | ./lautfm-streamer yourstation xxxx-xxxx-xxxx-xxxx metadata
```
… and in a separate process (using https://github.com/fabian1811/cmdline-musicdetection):
```
cat pipe | ./gn_stream_id client_id client_id_tag license_file metadata
```
Not the gracenote stream identifier will query the gracenote database to recognize the current song and update the metadata file with the appropriate metadata string. lautfm-streamer will catch the changes of the metadata file.

Requirements
------------

libshout must be installed.

Restrictions
-----

As cross compilation does't work for GO projects with C bindings I can only provide the linux 64 bit binary for now.
