use pest_consume::{Parser as PestParser};

#[derive(PestParser)]
#[grammar = "./grammar.pest"]
pub struct Parser;

#[pest_consume::parser]
impl Parser {

}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let successful_parse = Parser::parse(Rule::expression, "1 + 1");
        println!("{:?}", successful_parse);
    }
}
