local counter = 0

request = function()
    counter = counter + 1
    headers = {}
    headers["Content-Type"] = "application/json"
    body = '{"title": "Story ' .. counter .. '"}'
    return wrk.format("POST", "/todos/api/v1/stories", headers, body)
end
