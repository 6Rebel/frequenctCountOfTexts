# frequenctCountOfTexts
Create a service that accepts input as text and provides Json Output as Top ten most used words and times of occurrence in the Another Service.

# ENDPOINTS
POST API: /api/topTenWords

# Request: {
    "text": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum. Contrary to popular belief, Lorem Ipsum is not simply random text."
}


# Response: 
[{"word":"the","count":6},{"word":"Lorem","count":5},{"word":"of","count":4},{"word":"Ipsum","count":4},{"word":"and","count":3},{"word":"simply","count":2},{"word":"has","count":2},{"word":"text","count":2},{"word":"It","count":2},{"word":"with","count":2}]
