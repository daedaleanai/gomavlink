# gomavlink/C
MAVLink Encoding and Decoding in C

NOTE ON THE LICENSE: 
The generator code ../mavgen and the .tmpl files are under the MIT license in the root directory.
You can always run ../mavgen on whatever XML you feel running it on and use the output however you see fit, but 
the generated files here are derived from XML code under the LGPL, i include them here as a service, but they are under the LGPL.

This subdirectory of gomavlink contains templates for C header and code files for ../mavgen.
Using them, the include/ and src/ directories contain everything to serialize and deserialize
all dialects in the original mavlink XML definitions.

The generated code has no dependencies other than a handful of standard C include headers (stddef and stdint, string and assert), and may be used under the same license as the original XML files you 

The C code is split into a header include/${dialect}.h and separate source files src/${dialect}_enc.c, _dec.c and _crc.c used by both.

(_fmt.c, to print the messges in a debug format, is WIP)


the code is meant to be copied in-place to your project, you can always generate it in place by running the commands as exemplified in the update.sh script.





