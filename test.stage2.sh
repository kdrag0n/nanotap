#!/system/bin/sh -e
cd /data/local/tmp
rm -fr nt
mkdir nt
cd nt
mv ../config.toml ../nanotap .
echo 'Starting program'
./nanotap
