from fastapi import FastAPI
from fastapi.responses import HTMLResponse

from shared.database import Base, engine
from note import note_router


Base.metadata.create_all(engine)

app = FastAPI()
app.include_router(note_router)


@app.get("/", response_class=HTMLResponse, include_in_schema=False)
async def rapidoc():
    return f"""
        <!doctype html>
        <html>
            <head>
                <meta charset="utf-8">
                <script
                    type="module"
                    src="https://unpkg.com/rapidoc/dist/rapidoc-min.js"
                ></script>
            </head>
            <body>
                <rapi-doc spec-url="{app.openapi_url}"></rapi-doc>
            </body>
        </html>
    """
