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
  (fun nums -> [List.map int_of_string nums |> List.fold_left ( * ) 1])
  <$> 
  (string "mul(" *> k <* char ')')

let _read_until p =
  many_till any_char (p)

let drop_until p =
  fix (fun recurse ->
    p <|> any_char *> recurse
  )

let parser = many (drop_until mul)

let () =
  let input = "mul(2,4)%&mul[3,7]!@^do_not_mul(5,5)en(mul(11,8)mul(8,5))" in
  let result = parse_string ~consume:Prefix parser input in
  match result with
  (* | Ok res -> printf("%s") (res |> String.concat " ") *)
  | Ok res -> (printf "%d ") (res |> List.flatten |> List.fold_left ( + ) 0 )
  (* | Ok res -> (printf "%s") res *)
  | Error err -> printf("ERROR: %s") err


(* PART TWO *)
(* -------------------------------------------------------------------------- *)