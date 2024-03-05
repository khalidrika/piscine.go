#!/bin/bash
curl -s https://learn.zonie01oujda.ma/assets/superhero/all.json | jq -r ' .[]| select(.id == 170)| .name , .powerstats.power , .appearance.gender '
