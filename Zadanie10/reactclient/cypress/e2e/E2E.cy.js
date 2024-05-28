/* global cy */
describe('Testy funkcjonalne bez serwera', () => {
    beforeEach(()=>
    {
        cy.visit('http://localhost:3000')
    })
    it('Przejdź całą aplikację', () => {
     cy.title().should("eq", "React App")
     cy.contains("button", "Przejdź do produktów").click()
     cy.contains("div", "Produkty")
     cy.contains("button", "Przejdź do koszyka").click()
     cy.contains("div", "Twój koszyk:")
     cy.contains("button", "Przejdź do płatności").click()
     cy.contains("div", "Dokonaj płatności")
     cy.contains("button", "Wróć do koszyka").click()
     cy.contains("div", "Twój koszyk:")
     cy.contains("button", "Strona główna").click()
     cy.contains("div", "Reaction Nation")
    })
    it("Spróbuj dokonać płatności bez podania nazwy użytkownika", () =>{
        cy.contains("button", "Przejdź do koszyka").click()
        cy.contains("button", "Wyczyść koszyk").click()
        cy.contains("button", "Przejdź do płatności").click()
        cy.contains("button", RegExp("Kontynuuj", "g")).click()
        cy.on("window:alert", (str)=>{
            expect(str).to.equal("Please fill out this field.")
        })
        })
    it  ("Spróbuj dokonać płatności bez podania metody płatności", () =>{
        cy.contains("button", "Przejdź do koszyka").click()
        cy.contains("button", "Wyczyść koszyk").click()
        cy.contains("button", "Przejdź do płatności").click()
        cy.get(".Payments > form:nth-child(5) > input:nth-child(2)").type("Użytkownik")
        cy.contains("button", RegExp("Kontynuuj", "g")).click()
        cy.on("window:alert", (str)=>{
            expect(str).to.equal("Wybierz właściwą metodę płatności")
        })
    })
    })