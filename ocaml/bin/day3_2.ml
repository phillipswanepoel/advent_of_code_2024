open Angstrom
open Printf

let is_digit = function '0'..'9' -> true | _ -> false

let int3 =
  take_while1 is_digit
  >>= fun digits ->
  if String.length digits <= 3 then 
    return digits
  else 
    fail "number too large"

let mul =
  let k = sep_by1 (char ',') int3 in
  (string "mul(" *> k <* char ')')

let _read_until p =
  many_till any_char (p)

let drop_until p =
  fix (fun recurse ->
    p <|> any_char *> recurse
  )