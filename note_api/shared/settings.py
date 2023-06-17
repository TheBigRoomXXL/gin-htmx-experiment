from pydantic import BaseSettings


class Settings(BaseSettings):
    APP_NAME: str = "Notes API"
    DB_URI: str = 'sqlite:///:memory:'

settings = Settings()