#!/bin/sh

echo "{
  \"HttpPort\": \"${HttpPort}\",
  \"GinMode\": \"${GinMode}\",
  \"DependTools\": [\"npm\",\"yarn\"],
  \"LogMode\": ${LogMode},
  \"AddressUrl\" : \"${AddressUrl}\",
}" > /marewood/config.json
exec /marewood/marewood
