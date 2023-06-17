
from fastapi import APIRouter
from note.schema import NoteSchema
# from note.service import service_note as snote

router = APIRouter()

@router.post("/notes")
async def create_note(note: NoteSchema):
    # new_note = snote.create_one(note.dict())
    # return new_note
    return note
