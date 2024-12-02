open Printf

let parse_file filename =
  let ic = open_in filename in
  let rec read_lines acc =
    try
      let line = input_line ic in
      let numbers = List.map int_of_string (String.split_on_char ' ' line) in
      read_lines (numbers :: acc)
    with End_of_file ->
      close_in ic;
      List.rev acc
  in
  read_lines []

let input = parse_file "data/2/test"

(* PART ONE *)
let rec diffs = function
| [] | [_] -> []
| h1 :: (h2 :: _ as t) -> 
    h2 - h1 :: diffs t

let isValid lst = 
  (List.for_all (fun x -> x > 0) lst || List.for_all (fun x -> x < 0) lst)
  &&
  (List.for_all (fun x -> abs(x) >= 1 && abs(x) <= 3) lst)

let part_one =
input
|> List.map diffs 
|> List.map isValid
|> List.fold_left (fun acc x -> if x then acc + 1 else acc) 0 

let () = (printf "%d\n") part_one

(* PART TWO *)

(* let part_two =
  input
  |> List.map diffs
  |> List.map isValid
  |>  *)

 
