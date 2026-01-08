import axios from 'axios';
import { getConfig } from './config';

export class IndexerAPI {
  constructor() {
    const config = getConfig();
    this.baseURL = config.indexerApiUrl;
    this.client = axios.create({
      baseURL: `${this.baseURL}/api`,
      timeout: 10000,
    });
  }

  async getAllClues() {
    const response = await this.client.get('/clues');
    return response.data;
  }

  async getClue(clueId) {
    const response = await this.client.get(`/clues/${clueId}`);
    return response.data;
  }

  async getEvents(params = {}) {
    const response = await this.client.get('/events', { params });
    return response.data;
  }

  async getEventsByClueId(clueId, params = {}) {
    const response = await this.client.get('/events', {
      params: { clueId, ...params }
    });
    return response.data;
  }

  async getEventsByType(eventType, params = {}) {
    const response = await this.client.get('/events', {
      params: { type: eventType, ...params }
    });
    return response.data;
  }

  async getEventsByInitiator(initiator, params = {}) {
    const response = await this.client.get('/events', {
      params: { initiator, ...params }
    });
    return response.data;
  }

  async getOwnership(params = {}) {
    const response = await this.client.get('/ownership', { params });
    return response.data;
  }

  async getOwnershipByAddress(owner) {
    const response = await this.client.get('/ownership', {
      params: { owner }
    });
    return response.data;
  }

  async getOwnershipByClueId(clueId) {
    const response = await this.client.get('/ownership', {
      params: { clueId }
    });
    return response.data;
  }
}

export class GatewayAPI {
  constructor() {
    const config = getConfig();
    this.baseURL = config.gatewayApiUrl;
    this.client = axios.create({
      baseURL: this.baseURL,
      timeout: 10000,
    });
  }

  async getLinkedPublicKey(ethereumAddress) {
    try {
      const response = await this.client.get('/link', {
        params: { ethereumAddress }
      });
      return response.data;
    } catch (error) {
      if (error.response?.status === 404) {
        return { success: false, skavenge_public_key: null };
      }
      throw error;
    }
  }
}
