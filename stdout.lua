-- Implementations should not directly write to stdout but we should
-- save that somewere for logs etc.
function sum(a, b)
    print("Oh no I can print an obnoxious ad here.")
    return a + b
end
