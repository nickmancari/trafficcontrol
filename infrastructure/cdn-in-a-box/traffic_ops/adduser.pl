#!/usr/bin/env perl
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.

# adduser.pl creates the sql necessary to add a user to the database for traffic_ops login.
# Usage:   adduser.pl <username> <password> <role>
#  -- the password is encrypted appropriately to be compatible with Traffic Ops.
#
use strict;
use Crypt::ScryptKDF qw{ scrypt_hash };

if ($#ARGV < 2) {
    die "Usage: $ARGV[0] <username> <password> <role>\n";
}

my $username = shift // 'admin';
my $password = shift or die "Password is required\n";
my $role = shift // 'admin';
my $tenant = shift // 'root';

# Skip the insert if the admin 'username' is already there.
my $hashed_passwd = hash_pass( $password );
print <<"ADMIN";
insert into tm_user (username, role, local_passwd, confirm_local_passwd, tenant_id)
    values  ('$username',
            (select id from role where name = '$role'),
            '$hashed_passwd',
            '$hashed_passwd',
            (SELECT id FROM tenant WHERE name='$tenant'))
    ON CONFLICT (username) DO NOTHING;
ADMIN

sub hash_pass {
    my $pass = shift;
    return scrypt_hash($pass, \64, 16384, 8, 1, 64);
}
