#!/system/bin/sh -e
cd /data/local/tmp
rm -fr nt
mkdir nt
cd nt
cp ../config.toml ../nanotap .
echo 'Starting program'
./nanotap
