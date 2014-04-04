#!/usr/bin/env python
# -*- coding: utf-8 -*-

from __future__ import absolute_import
from __future__ import division
from __future__ import print_function
from __future__ import unicode_literals

from future.builtins import *

import fileinput
import json
import sys
import time

import onetimepass as otp

def main():
    secrets_data = sys.stdin.read()
    secrets = json.loads(secrets_data)

    default_color='\033[39m'

    while True:
        print('\033[2J\033[1;1H\033[39m\033[49m', end='')
        timeleft = 30 - int(time.time()) % 30
        print("Time left: ", timeleft)

        if timeleft < 5:
            color = '\033[31m'
        elif timeleft < 10:
            color = '\033[35m'
        else:
            color = '\033[34m'

        for desc, key in secrets:
            token = otp.get_totp(key.replace(' ', '').encode())
            print("%s%06d%s: (%s)" % (color, token, default_color, desc))
        time.sleep(1)

if __name__ == '__main__':
    sys.exit(main())

__all__ = []
