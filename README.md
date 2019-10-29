# HomeAutomationv2.0
Setup Dockers
docker run -d --name behomeserver --restart always --volume switch:/app sarafdarpundlik/homeappli:pi3
docker run -d -p 80:80 --name ui --link=behomeserver:behomeserver homeui
