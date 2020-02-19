#!/usr/bin/env perl

use Cache::Memcached::Fast;

$memc = Cache::Memcached::Fast->new({
    'servers' => ['localhost:11211']
});

$x = 100;

$memc->set("xxx", $x);

print $memc->get("xxx");

1;
