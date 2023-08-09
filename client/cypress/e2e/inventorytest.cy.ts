describe('inventorytest', () => {
      it('passes', () => {
        cy.visit('localhost:3000')
        /* ==== Generated with Cypress Studio ==== */
        cy.get('[style="margin-right: 12px; cursor: none;"]').click();
        cy.get('[placeholder="Username"]').clear('oi');
        cy.get('[placeholder="Username"]').type('oidc');
        cy.get('[placeholder="Password"]').clear();
        cy.get('[placeholder="Password"]').type('password');
        cy.get('.btn-primary').click();
        /* ==== Generated with Cypress Studio ==== */
        cy.get('[d="M12 21.6099V12.5399"]').click();
        /* ==== End Cypress Studio ==== */
        /* ==== Generated with Cypress Studio ==== */
        cy.get('[d="M12 21.6099V12.5399"]').click();
        /* ==== End Cypress Studio ==== */
      })
})