open Angstrom
open Printf

(* PARSING 
yes this is overkill and weird, but trying to learn angstrom *)
let is_space = function
  | ' ' | '\t' -> true
  | _ -> false

let whitespace = take_while is_space

let number = take_while1 (function '0'..'9' -> true | _ -> false)
             >>| int_of_string

let line = 
  let* left = number in
  let* _ = whitespace in
  let* right = number in
  let* _ = end_of_line in
  return (left, right)

let parser = Angstrom.many line

let read_file filename =
  let in_channel = In_channel.open_text filename in
  let contents = In_channel.input_all in_channel in
  In_channel.close in_channel;
  contents

let parse_input filename =
  let input_string = read_file filename in
  match Angstrom.parse_string ~consume:Angstrom.Consume.All parser input_string with
  | Ok pairs -> 
    let lefts, rights = List.split pairs in
    (lefts, rights)
  | Error msg -> failwith msg

let data = parse_input "data/test"
let first_sorted = fst data |> List.sort compare
let second_sorted = snd data |> List.sort compare
let pairs = List.combine first_sorted second_sorted 

(* PART ONE *)
let output1 = 
 pairs
 |> List.map (fun (x, y) -> abs(x-y)) 
 |> List.fold_left ( + ) 0

let () = (printf "%d\n") output1

(* PART TWO *)
(*encode counts of characters in second list*)
(* [1;3;3] -> [(1, 1); (2; 3)]*)

(* Encoding counts since I had this function lying around*)
let encode list =
  let rec aux count = function
    | [] -> []
    | [x] -> (count+1, x) :: []
    | a :: (b :: _ as t) ->
      if a = b then aux (count+1) t
      else (count+1, a) :: aux 0 t
  in aux 0 list

let second_encoded = encode second_sorted
let folder h = List.fold_left (fun acc (x, y) -> if y = h then acc + (x * y) else acc) 0 second_encoded
let output2 = List.map (fun h -> folder h) first_sorted |> List.fold_left ( + ) 0

let () = (printf "%d\n") output2

(* let first_filtered = List.filter (fun a -> List.mem a second_sorted) first_sorted
let first_encoded = encode first_filtered
let second_filtered = List.filter (fun a -> List.mem a first_filtered ) second_sorted
let second_encoded = encode second_filtered
let output2 =
  List.fold_left2 (fun acc (x, _) (a, b) -> acc + (a * b * x)) 0 first_encoded second_encoded 

let () = (printf "%d\n") output2  *)