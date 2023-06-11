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
def outputData(name, created_at, state):
    if created_at >= 9 and created_at <= 17:
        print('%s\t%s'%(name, state))

# mapper function
def mapper(tags=TAGS):    
               
    for sample in sys.stdin:
        data = compile(sample)
        
        created_at = 0
        state = 'others'
        lat = None
        lng = None
        # extract wanted data
        text = data[2]
        # using lower case to make it case insensitive 
        try:
            created_at = int(data[0].split(" ")[1].split(":")[0])
        except:
            created_at = 0
        try:
            lat = float(data[13])
            lng = float(data[14])
        except:
            lat = None
            lng = None
            
        if not lat or not lng:
            continue
            
        if 40.4772 <= lat and lat <= 45.0153 and -79.7624 <= lng and lng <= -71.7517:
            state = 'new york'
        elif 32.5121 <= lat and lat <= 42.0126 and -124.6509 <= lng and lng <= -114.1315:
            state = 'california'
        else:
            continue
            
        # find related operation then operate
        found_biden = any(term in text for term in tags['Biden'])
        found_trump = any(term in text for term in tags['Trump'])

        if found_biden and not found_trump:
            outputData('Biden', created_at, state)
        elif found_trump and not found_biden:
            outputData('Trump', created_at, state)
        elif found_biden and found_trump:
            outputData('Both', created_at, state)
                
mapper(TAGS)
