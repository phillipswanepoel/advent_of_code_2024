open Angstrom
open Printf

let test = 
  "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

let is_digit = function '0'..'9' -> true | _ -> false

let int3 =
  take_while1 (function 
    | '0'..'9' -> true 
    | _ -> false)
  >>= fun digits ->
  if String.length digits <= 3 then 
    return digits
  else 
    fail "number too large"

let mul =
  string "mul("
  *>                
  int3 
  >>= fun first ->      
  char ','
  *>                   
  int3 
  >>= fun second ->     
  char ')' 
  >>|               
  fun _closing -> int_of_string first * int_of_string second

let any = peek_char >>= function
  | Some _ -> advance 1 >>| fun () -> -1
  | None -> fail "end of input"

let find_open = mul <|> any
let find_opens = many (find_open)

let test = 
  "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
let remove_fail = List.filter (fun x -> x <> -1)


let ans input = Angstrom.parse_string ~consume:Prefix find_opens input
|> Result.get_ok 
|> remove_fail
|> List.fold_left (+) 0


let read_file filename =
  let in_channel = In_channel.open_text filename in
  let contents = In_channel.input_all in_channel in
  In_channel.close in_channel;
  contents
let input_string = read_file "data/3/input"
let output1 = ans input_string

let () = (printf "%d\n") output1

(* PART TWO *)
(* -------------------------------------------------------------------------- *)

let fail_if_dont = 
  peek_string 7
  >>= function
  | "don't()" -> fail "don't"
  | _ -> take 1

let fail_if_do = 
  peek_string 4
  >>= function
  | "do()" -> fail "do"
  | _ -> take 1

let dos_and_donts =
  (peek_string 7
  >>= function
  | "don't()" -> take 7
  | s when String.starts_with ~prefix:"do()" s -> take 4
  | _ -> take 1) <|> take 1

let ans2 = Angstrom.parse_string ~consume:Prefix (many dos_and_donts) input_string
let k = (ans2 |> Result.get_ok)

let process lst = 
  let rec aux switch lst = 
  match lst with
  | [] -> []
  | h :: t -> 
    if h = "don't()" then aux false t
    else if h = "do()" then aux true t
    else (if switch then h :: aux switch t else aux switch t)
  in aux true lst

let processed = process k |> String.concat ""
let ans2 = ans processed
