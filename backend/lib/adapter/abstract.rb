# frozen_string_literal: true

module Adapter
  class Abstract
    def balance; end

    def order_book
      raise 'Override method in children'
    end

    def fee
      raise 'Override method in children'
    end

    def wallet_status
      raise 'Override method in children'
    end

    def create_order
      raise 'Override method in children'
    end

    def cancel_order
      raise 'Override method in children'
    end

    def open_orders
      raise 'Override method in children'
    end

    def order_status
      raise 'Override method in children'
    end

    def market
      raise 'Override method in children'
    end
  end
end
