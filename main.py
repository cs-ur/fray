import discord
from dotenv import dotenv_values

intents = discord.Intents.default()
intents.message_content = True
dotenv_config = dotenv_values('.env')

discord_token = dotenv_config.get('DISCORD_TOKEN')
prefix = dotenv_config.get('PREFIX', '!')

if discord_token is None or prefix is None:
    print("Discord token or prefix not found in .env file")
    exit(1)

client = discord.Client(intents=intents)

@client.event
async def on_ready():
    print(f"Logged in as {client.user}")

@client.event
async def on_message(message):
    if message.author == client.user: return
    
    if message.content == "!ping":
        await message.channel.send("Pong!")

client.run(token=discord_token)

