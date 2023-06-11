docker compose up -d
docker ps
docker exec -it <mycontainer> bash
docker exec -it namenode bash


hdfs dfs -mkdir -p user/root/input
hdfs dfs -copyFromLocal dataset.csv user/root/input


hdfs dfs -ls user/root/input


hadoop jar /opt/hadoop-3.2.1/share/hadoop/tools/lib/hadoop-streaming-3.2.1.jar  -file /home/app1/mapper.py -mapper "python3 mapper.py" -file /home/app1/reducer.py -reducer "python3 reducer.py" -input /user/root/input/dataset.csv -output /user/root/output1

hadoop jar /opt/hadoop-3.2.1/share/hadoop/tools/lib/hadoop-streaming-3.2.1.jar  -file /home/app2/mapper.py -mapper "python3 mapper.py" -file /home/app2/reducer.py -reducer "python3 reducer.py" -input /user/root/input/dataset.csv -output /user/root/output2

hadoop jar /opt/hadoop-3.2.1/share/hadoop/tools/lib/hadoop-streaming-3.2.1.jar  -file /home/app3/mapper.py -mapper "python3 mapper.py" -file /home/app3/reducer.py -reducer "python3 reducer.py" -input /user/root/input/dataset.csv -output /user/root/output3
