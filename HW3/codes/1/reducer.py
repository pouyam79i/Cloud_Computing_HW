#!/usr/bin python3
import sys


# result holder
RESULT = {
    "Biden": {
        'likes': 0,
        'retweets': 0,
        'Twitter Web App': 0,
        'Twitter for iPhone': 0,
        'Twitter for Android': 0
    },
    "Trump": {
        'likes': 0,
        'retweets': 0,
        'Twitter Web App': 0,
        'Twitter for iPhone': 0,
        'Twitter for Android': 0
    },
    "Both": {
        'likes': 0,
        'retweets': 0,
        'Twitter Web App': 0,
        'Twitter for iPhone': 0,
        'Twitter for Android': 0
    }
}


# data reducer 
def reducer():
    for line in sys.stdin:
        line = line.strip()
        name, likes, retweets, TFW, TFI, TFA = line.split('\t')
        try:
            likes = int(likes)
            retweets = int(retweets)
            TFW = int(TFW)
            TFI = int(TFI)
            TFA = int(TFA)
            RESULT[name]['likes'] += likes
            RESULT[name]['retweets'] += retweets
            RESULT[name]['Twitter Web App'] += TFW
            RESULT[name]['Twitter for iPhone'] += TFI
            RESULT[name]['Twitter for Android'] += TFA
        except:
            pass


# output results
def outputResult():
    for name in RESULT:
        likes = RESULT[name]['likes']
        retweets = RESULT[name]['retweets']
        TFW = RESULT[name]['Twitter Web App']
        TFI = RESULT[name]['Twitter for iPhone']
        TFA = RESULT[name]['Twitter for Android']
        print('%s\t%s\t%s\t%s\t%s\t%s'%(name, likes, retweets, TFW, TFI, TFA))


# run app:
reducer()
outputResult()
