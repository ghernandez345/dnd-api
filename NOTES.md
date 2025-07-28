api structure

/api/:year/:resource?query_string_filters

examples:

api/2014/monsters?name="zombie" -> get the zombie monster
api/2024/classes?name="ranger" -> get the ranger class
api/2014/magic_items?rarity="legendary" -> all magic items that are legendary rarity
api/2014/monsters?ac=">=13"&cr="<4" -> get all monsters whos armour class is greater than or equal to 13 and challenge rating is less then 4
