#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
Easily extract iksm_session key by connecting Ika-Ring2 through mitmproxy.

Usage: 1. Run `mitmdump -p PORT -s getIksmSession.py` (PORT is something like 8080).
       2. Connect the mobile phone to the proxy server.
       3. Connect to the Ika-Ring2 from Nintendo-Switch app.
       4. iksm_session key should be printed on the screen.
"""

def response(flow):
    if "https://app.splatoon2.nintendo.net/" in flow.request.url:
        headers = flow.request.headers
        if headers.__contains__('cookie'):
            cookie = flow.request.headers['cookie']
            d = {}
            for item in cookie.split(';'):
                key, val = item.strip().split('=')
                d.update({key: val})
            if d.__contains__('iksm_session'):
                print('-----------------------------')
                print('Your iksm_session is:')
                print(d['iksm_session'])
                print('-----------------------------')
