# frozen_string_literal: true

class Exchange
  # All exchange's methods to adapter
  delegate :get_orders, to: :adapter

  # @param adapter [Adapter]
  # @
  def initialize(adapter)
    @adapter = adapter
  end
end
