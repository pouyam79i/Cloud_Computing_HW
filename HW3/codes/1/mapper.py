#!/usr/bin python3
import sys
import re

# DEFAULT:
# using this pattern to compile data using regex
RE_PATTERN_C = re.compile(r''',(?=(?:(?:[^"]*"){2})*[^"]*$)''')
# tags filter
TAGS = {
    'Biden':('#JoeBiden', '#Biden'),
    'Trump':('#DonaldTrump', '#Trump')
}

# compiles raw data
def compile(raw_data):
    data = raw_data.strip()                                                
    data = RE_PATTERN_C.split(data)
    # now we have index=2 as text, index=3 as likes, index=4 as retweets, index=5 as source
    return data

# prints output data
def outputData(name, likes, retweets, source):
    TWA = 0
    TFI = 0
    TFA = 0
    if source == 'Twitter Web App':
        TWA = 1
    elif source == 'Twitter for iPhone':
        TFI = 1
    elif source == 'Twitter for Android':
        TFA = 1
    print('%s\t%s\t%s\t%s\t%s\t%s'%(name, likes, retweets, TWA, TFI, TFA))

# mapper function
def mapper(tags=TAGS):
    for sample in sys.stdin:
        data = compile(sample)
        
        likes = 0
        retweets = 0
        source = 'null'
        # extract wanted data
        text = data[2]
        try:
            likes = int(float(data[3]))
        except:
            likes = 0
        try:
            retweets = int(float(data[4]))
        except:
            retweets = 0
        try:
            source = data[5]
        except:
            source = 'null'

        # find related operation then operate
        found_biden = any(term in text for term in tags['Biden'])
        found_trump = any(term in text for term in tags['Trump'])

        if found_biden and not found_trump:
            outputData('Biden', likes, retweets, source)
        elif found_trump and not found_biden:
            outputData('Trump', likes, retweets, source)
        elif found_biden and found_trump:
            outputData('Both', likes, retweets, source)  
    
mapper(TAGS)
