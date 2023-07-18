const express = require('express')

function getRequestPayload(req) {
  const payload = req.body
  if (!payload.headers) {
    payload.headers = {}
  }

  return payload
}

function stringifiedValidJson(res, body) {
  return res.json({
    ...body,
    control: { break: 401 },
    body: JSON.stringify({
      errors: [
        {
          message: 'Not authenticated.',
          extensions: {
            code: 'ERR_UNAUTHENTICATED',
          },
        },
      ],
    })
  })
}

function handleValidStringifiedJson(req, res) {
  const payload = getRequestPayload(req)
  return stringifiedValidJson(res, payload)
}

function validJson(res, body) {
  return res.json({
    ...body,
    control: { break: 401 },
    body: {
      errors: [
        {
          message: 'Not authenticated.',
          extensions: {
            code: 'ERR_UNAUTHENTICATED',
          },
        },
      ],
    }
  })
}

function handleValidJson(req, res) {
  const payload = getRequestPayload(req)
  return validJson(res, payload)
}

const port = process.env.PORT || 3000
const app = express()
app.use(express.json())
app.post('/stringified-valid-json', handleValidStringifiedJson)
app.post('/valid-json', handleValidJson)
app.listen(port, () => {
  console.log(`🚀 Coprocessor running on port ${port}`)
})
