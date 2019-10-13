/*
 * Climate Hero
 * A web app to incentivize users to be more eco-friendly
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiException;
import org.openapitools.client.model.CheckReturn;
import org.openapitools.client.model.Error;
import org.openapitools.client.model.InlineObject;
import org.openapitools.client.model.InlineObject1;
import org.openapitools.client.model.InlineObject2;
import org.openapitools.client.model.Name;
import org.openapitools.client.model.NewUser;
import org.openapitools.client.model.UserInfo;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for DefaultApi
 */
@Ignore
public class DefaultApiTest {

    private final DefaultApi api = new DefaultApi();

    
    /**
     * 
     *
     * Creates a new Climate Hero user 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addUserTest() throws ApiException {
        NewUser newUser = null;
        Name response = api.addUser(newUser);

        // TODO: test validations
    }
    
    /**
     * 
     *
     * Checks an item
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void checkItemTest() throws ApiException {
        InlineObject1 inlineObject1 = null;
        CheckReturn response = api.checkItem(inlineObject1);

        // TODO: test validations
    }
    
    /**
     * 
     *
     * Gets user info
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getUserInfoTest() throws ApiException {
        String email = null;
        UserInfo response = api.getUserInfo(email);

        // TODO: test validations
    }
    
    /**
     * 
     *
     * Logs in a user
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void loginTest() throws ApiException {
        InlineObject inlineObject = null;
        Name response = api.login(inlineObject);

        // TODO: test validations
    }
    
    /**
     * 
     *
     * updates a user list
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateListTest() throws ApiException {
        InlineObject2 inlineObject2 = null;
        api.updateList(inlineObject2);

        // TODO: test validations
    }
    
}
