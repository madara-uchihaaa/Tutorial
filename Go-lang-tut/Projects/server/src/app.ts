import 'dotenv/config';
import express from 'express';
import bodyParser from 'body-parser';
import { OpenAI } from 'openai'


const app = express()
app.use(bodyParser.json())

const openai = new OpenAI({
    apiKey: process.env.CHAT_GPT_API,
})


app.post('/ask', async (req, res) => {
    const { question } = req.body;
    console.log('Question', question);
    try {
        const chatComplete = await openai.chat.completions.create({
            messages: [{
                role: 'user',
                content: question,
            }],
            model: 'gpt-3.5-turbo',
        });
        console.log('Chat complete', chatComplete.choices[0].message.content);
        res.send({
            response: chatComplete.choices[0].message.content,
            statusCode: 200,
        });
    } catch (error) {
        console.log('Error while asking question', error);
        res.send({
            response: 'Error while asking question',
            statusCode: 500,
        });
    }
})

app.listen(process.env.PORT, () =>
    console.log(`Example app listening on port ${process.env.PORT}!`),
)

