describe('signup', () => {
  it('passes', () => {
    cy.visit('localhost:3000')
    /* ==== Generated with Cypress Studio ==== */
    cy.get('.btn-secondary').click();
    cy.get('[placeholder="Username"]').clear('oi');
    cy.get('[placeholder="Username"]').type('oidc');
    cy.get('[placeholder="Email"]').clear();
    cy.get('[placeholder="Email"]').type('oidc@gmail.com');
    cy.get('[placeholder="Password"]').clear();
    cy.get('[placeholder="Password"]').type('password');
    cy.get('[placeholder="Password Check"]').clear();
    cy.get('[placeholder="Password Check"]').type('password');
    cy.get('.signupbutton > .btn').click();
    /* ==== End Cypress Studio ==== */
    /* ==== Generated with Cypress Studio ==== */
    cy.get('.btn-secondary').click();
    cy.get('[placeholder="Username"]').clear('oi');
    cy.get('[placeholder="Username"]').type('oidc');
    cy.get('[placeholder="Email"]').clear();
    cy.get('[placeholder="Email"]').type('oidc@gmail.com');
    cy.get('[placeholder="Password"]').clear();
    cy.get('[placeholder="Password"]').type('password');
    cy.get('[placeholder="Password Check"]').clear();
    cy.get('[placeholder="Password Check"]').type('password');
    cy.get('.signupbutton > .btn').click();
    /* ==== End Cypress Studio ==== */
  })
})