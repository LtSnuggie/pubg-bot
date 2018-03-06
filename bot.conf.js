{
  "mysql" : {
    "user": "pubgbot",
    "pass": "Pass$1",
    "database": "PUBG",
    "tables": [
      {
        "name": "users",
        "col_names": ["discordId","gamertag"],
        "col_types": ["text","text"]
      },
      {
        "name": "aliases",
        "col_names": ["alias","gamertag"],
        "col_types": ["text","text"]
      },
      {
        "name": "watch",
        "col_names": ["gamertag"],
        "col_types": ["text"]
      },
      {
        "name": "stats",
        "key": "id",
        "col_names": ["id",
                      "gamertag",
                      "Matches Played",
                      "Matches Won",
                      "Win Percentage",
                      "Total Kills",
                      "KDA",
                      "Total Headshots",
                      "Headshot Percentage",
                      "Solo Matches Played",
                      "Solo Wins",
                      "Solo Win Percentage",
                      "Duo Wins",
                      "Squad Wins",
                      "Days",
                      "Hours",
                      "Minutes"
                    ],
        "col_types": ["int",
                      "text",
                      "float",
                      "float",
                      "float",
                      "float",
                      "float",
                      "float",
                      "float",
                      "float",
                      "float",
                      "float",
                      "float",
                      "float",
                      "float",
                      "float",
                      "float"
                    ]
      }
    ]
  },
  "test_channel" : "415702469390106635",
  "log_channel" : "415702469390106635",
  "token" : "NDEyMzM5Nzg3NjQzNjE3Mjgw.DW5ieQ.2R2nSvhuxlRYfFHfUjk9LNgEKQo",
  "poll_interval" : 1
}
