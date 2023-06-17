from pydantic import BaseModel, Field

class Note(BaseModel):
    title:str = Field(
        title="The title of the note", max_length=255
    )
    body:str = Field(
        title="The body of the note"
    )
    # tags
    # links