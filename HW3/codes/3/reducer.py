#!/usr/bin python3
import sys

# result holder
RESULT = {
    'new york': {
        'Biden': 0,
        'Trump': 0,
        'Both': 0
    },
    'california': {
        'Biden': 0,
        'Trump': 0,
        'Both': 0
    }
}

# data reducer 
def reducer():
    for line in sys.stdin:
        line = line.strip()
        name, state = line.split('\t')
        try:
            RESULT[state][name] = RESULT[state][name] + 1
        except:
            pass
            

# output results
def outputResult():
    for state in RESULT:
        biden = RESULT[state]['Biden']
        trump = RESULT[state]['Trump']
        both = RESULT[state]['Both']
        total = biden + trump + both
        if total == 0:
            total = 1
        both = both / total
        biden = biden / total
        trump = trump / total
        print('%s\t%s\t%s\t%s\t%s'%(state, both, biden, trump, total))

# run app:
reducer()
outputResult()
