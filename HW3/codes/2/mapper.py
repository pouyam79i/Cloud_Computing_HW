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
STATES = ['new york', 'texas', 'california', 'florida']

# compiles raw data
def compile(raw_data):
    data = raw_data.strip()                                                
    data = RE_PATTERN_C.split(data)
    # now we have index=2 as text, index=3 as likes, index=4 as retweets, index=5 as source
    return data

# prints output data
def outputData(name, created_at, state):
    if state not in STATES:
        return
    if created_at >= 9 and created_at <= 17:
        print('%s\t%s'%(name, state))

# mapper function
def mapper(tags=TAGS):    
               
    for sample in sys.stdin:
        data = compile(sample)
        
        created_at = 0
        state = 'others'
        # extract wanted data
        text = data[2]
        # using lower case to make it case insensitive 
        try:
            state = data[18].lower()
        except:
            state = 'others'
        try:
            created_at = int(data[0].split(" ")[1].split(":")[0])
        except:
            created_at = 0
        
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
